package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ory/dockertest/v3/docker"

	oracletypes "github.com/umee-network/umee/v4/x/oracle/types"
	"github.com/umee-network/umee/v4/x/uibc"
)

func (s *IntegrationTestSuite) connectIBCChains() {
	s.T().Logf("connecting %s and %s chains via IBC", s.chain.id, gaiaChainID)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	exec, err := s.dkrPool.Client.CreateExec(docker.CreateExecOptions{
		Context:      ctx,
		AttachStdout: true,
		AttachStderr: true,
		Container:    s.hermesResource.Container.ID,
		User:         "root",
		Cmd: []string{
			"hermes",
			"create",
			"channel",
			s.chain.id,
			gaiaChainID,
			"--port-a=transfer",
			"--port-b=transfer",
		},
	})
	s.Require().NoError(err)

	var (
		outBuf bytes.Buffer
		errBuf bytes.Buffer
	)

	err = s.dkrPool.Client.StartExec(exec.ID, docker.StartExecOptions{
		Context:      ctx,
		Detach:       false,
		OutputStream: &outBuf,
		ErrorStream:  &errBuf,
	})
	s.Require().NoErrorf(
		err,
		"failed connect chains; stdout: %s, stderr: %s", outBuf.String(), errBuf.String(),
	)

	s.Require().Containsf(
		errBuf.String(),
		"successfully opened init channel",
		"failed to connect chains via IBC: %s", errBuf.String(),
	)

	s.T().Logf("connected %s and %s chains via IBC", s.chain.id, gaiaChainID)
}

func (s *IntegrationTestSuite) sendIBC(srcChainID, dstChainID, recipient string, token sdk.Coin) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	s.T().Logf("sending %s from %s to %s (%s)", token, srcChainID, dstChainID, recipient)
	cmd := []string{
		"hermes",
		"tx",
		"raw",
		"ft-transfer",
		dstChainID,
		srcChainID,
		"transfer",  // source chain port ID
		"channel-0", // since only one connection/channel exists, assume 0
		token.Amount.String(),
		fmt.Sprintf("--denom=%s", token.Denom),
		"--timeout-height-offset=1000",
	}

	if len(recipient) != 0 {
		cmd = append(cmd, fmt.Sprintf("--receiver=%s", recipient))
	}

	exec, err := s.dkrPool.Client.CreateExec(docker.CreateExecOptions{
		Context:      ctx,
		AttachStdout: true,
		AttachStderr: true,
		Container:    s.hermesResource.Container.ID,
		Cmd:          cmd,
	})
	s.Require().NoError(err)

	var (
		outBuf bytes.Buffer
		errBuf bytes.Buffer
	)

	err = s.dkrPool.Client.StartExec(exec.ID, docker.StartExecOptions{
		Context:      ctx,
		Detach:       false,
		OutputStream: &outBuf,
		ErrorStream:  &errBuf,
	})

	s.Require().NoErrorf(
		err,
		"failed to send IBC tokens; stdout: %s, stderr: %s", outBuf.String(), errBuf.String(),
	)
	s.T().Log("successfully sent IBC tokens")
	s.T().Log("Waiting for 12 seconds to make sure trasaction is processed or include in the block")
	time.Sleep(time.Second * 12)
}

func queryUmeeTx(endpoint, txHash string) error {
	resp, err := http.Get(fmt.Sprintf("%s/cosmos/tx/v1beta1/txs/%s", endpoint, txHash))
	if err != nil {
		return fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("tx query returned non-200 status: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	txResp := result["tx_response"].(map[string]interface{})
	if v := txResp["code"]; v.(float64) != 0 {
		return fmt.Errorf("tx %s failed with status code %v", txHash, v)
	}

	return nil
}

func queryUmeeAllBalances(endpoint, addr string) (sdk.Coins, error) {
	resp, err := http.Get(fmt.Sprintf("%s/cosmos/bank/v1beta1/balances/%s", endpoint, addr))
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var balancesResp banktypes.QueryAllBalancesResponse
	if err := cdc.UnmarshalJSON(bz, &balancesResp); err != nil {
		return nil, err
	}

	return balancesResp.Balances, nil
}

func queryTotalSupply(endpoint string) (sdk.Coins, error) {
	resp, err := http.Get(fmt.Sprintf("%s/cosmos/bank/v1beta1/supply", endpoint))
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var balancesResp banktypes.QueryTotalSupplyResponse
	if err := cdc.UnmarshalJSON(bz, &balancesResp); err != nil {
		return nil, err
	}

	return balancesResp.Supply, nil
}

func queryExchangeRate(endpoint, denom string) (sdk.DecCoins, error) {
	resp, err := http.Get(fmt.Sprintf("%s/umee/oracle/v1/denoms/exchange_rates/%s", endpoint, denom))
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var exchangeRatesResponse oracletypes.QueryExchangeRatesResponse
	if err := cdc.UnmarshalJSON(bz, &exchangeRatesResponse); err != nil {
		return nil, err
	}

	return exchangeRatesResponse.ExchangeRates, nil
}

func queryHistroAvgPrice(endpoint, denom string) (sdk.Dec, error) {
	url := fmt.Sprintf("%s/umee/historacle/v1/avg_price/%s", endpoint, strings.ToUpper(denom))
	resp, err := http.Get(url)
	if err != nil {
		return sdk.Dec{}, fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		return sdk.Dec{}, err
	}

	var avgPriceResponse oracletypes.QueryAvgPriceResponse
	if err := cdc.UnmarshalJSON(bz, &avgPriceResponse); err != nil {
		return sdk.Dec{}, err
	}

	return avgPriceResponse.Price, nil
}

func queryOutflows(endpoint, denom string) (sdk.DecCoins, error) {
	resp, err := http.Get(fmt.Sprintf("%s/umee/uibc/v1/outflows?denom=%s", endpoint, denom))
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var outflowsResponse uibc.QueryOutflowsResponse
	if err := cdc.UnmarshalJSON(bz, &outflowsResponse); err != nil {
		return nil, err
	}

	return outflowsResponse.Outflows, nil
}

func queryUmeeDenomBalance(endpoint, addr, denom string) (sdk.Coin, error) {
	var zeroCoin sdk.Coin

	path := fmt.Sprintf(
		"%s/cosmos/bank/v1beta1/balances/%s/by_denom?denom=%s",
		endpoint, addr, denom,
	)
	resp, err := http.Get(path)
	if err != nil {
		return zeroCoin, fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		return zeroCoin, err
	}

	var balanceResp banktypes.QueryBalanceResponse
	if err := cdc.UnmarshalJSON(bz, &balanceResp); err != nil {
		return zeroCoin, err
	}

	return *balanceResp.Balance, nil
}

func (s *IntegrationTestSuite) queryUmeeBalance(
	umeeValIdx int,
	umeeTokenDenom string,
) (umeeBalance sdk.Coin, umeeAddr string) {
	umeeEndpoint := fmt.Sprintf("http://%s", s.valResources[umeeValIdx].GetHostPort("1317/tcp"))
	umeeAddress, err := s.chain.validators[umeeValIdx].keyInfo.GetAddress()
	s.Require().NoError(err)
	umeeAddr = umeeAddress.String()

	umeeBalance, err = queryUmeeDenomBalance(umeeEndpoint, umeeAddr, umeeTokenDenom)
	s.Require().NoError(err)
	s.T().Logf(
		"Umee Balance of tokens validator; index: %d, addr: %s, amount: %s, denom: %s",
		umeeValIdx, umeeAddr, umeeBalance.String(), umeeTokenDenom,
	)

	return umeeBalance, umeeAddr
}
