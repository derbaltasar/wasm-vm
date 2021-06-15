package dex

import (
	"errors"
	"fmt"
	"math/big"
)

func (pfe *fuzzDexExecutor) swapFixedInput(user string, swapPair SwapPair, amountA int,
	amountB int, statistics *eventsStatistics) error {

	tokenABefore, err := pfe.getTokens(user, swapPair.firstToken)
	if err != nil {
		return err
	}
	tokenBBefore, err := pfe.getTokens(user, swapPair.secondToken)
	if err != nil {
		return err
	}

	amountOutRaw, amountOutErr := pfe.querySingleResultStringAddr(pfe.ownerAddress, swapPair.address,
		"getAmountOut", fmt.Sprintf("\"str:%s\", \"%d\"", swapPair.firstToken, amountA))

	output, err := pfe.executeTxStep(fmt.Sprintf(`
	{
		"step": "scCall",
		"txId": "swap-fixed-input",
		"tx": {
			"from": "%s",
			"to": "%s",
			"value": "0",
			"function": "swapTokensFixedInput",
			"esdt": {
				"tokenIdentifier": "str:%s",
				"value": "%d"
			},
			"arguments": [
				"str:%s",
				"%d"
			],
			"gasLimit": "100,000,000",
			"gasPrice": "0"
		}
	}`,
		user,
		swapPair.address,
		swapPair.firstToken,
		amountA,
		swapPair.secondToken,
		amountB,
	))

	if err == nil {
		if output == nil {
			return errors.New("output is nil")
		}
		statistics.swapFixedInputHits += 1

		tokenAAfter, err := pfe.getTokens(user, swapPair.firstToken)
		if err != nil {
			return err
		}
		tokenBAfter, err := pfe.getTokens(user, swapPair.secondToken)
		if err != nil {
			return err
		}

		if amountOutErr == nil {
			if tokenAAfter.Cmp(big.NewInt(0).Sub(tokenABefore, big.NewInt(int64(amountA)))) != 0 {
				return errors.New("SWAP fixed input wrong amount A")
			}
			if tokenBAfter.Cmp(big.NewInt(0).Add(tokenBBefore, big.NewInt(0).SetBytes(amountOutRaw[0]))) != 0 {
				return errors.New("SWAP fixed input wrong amount B")
			}
		}
	} else {
		statistics.swapFixedInputMisses += 1

		pfe.log("swapFixedInput %s -> %s", swapPair.firstToken, swapPair.secondToken)
		pfe.log("could not swap because %s", output.ReturnMessage)
	}

	return nil
}

func (pfe *fuzzDexExecutor) swapFixedOutput(user string, swapPair SwapPair, amountA int,
	amountB int, statistics *eventsStatistics) error {

	tokenABefore, err := pfe.getTokens(user, swapPair.firstToken)
	if err != nil {
		return err
	}
	tokenBBefore, err := pfe.getTokens(user, swapPair.secondToken)
	if err != nil {
		return err
	}

	amountInRaw, amountInErr := pfe.querySingleResultStringAddr(pfe.ownerAddress, swapPair.address,
		"getAmountIn", fmt.Sprintf("\"str:%s\", \"%d\"", swapPair.secondToken, amountB))

	output, err := pfe.executeTxStep(fmt.Sprintf(`
	{
		"step": "scCall",
		"txId": "swap-fixed-input",
		"tx": {
			"from": "%s",
			"to": "%s",
			"value": "0",
			"function": "swapTokensFixedOutput",
			"esdt": {
				"tokenIdentifier": "str:%s",
				"value": "%d"
			},
			"arguments": [
				"str:%s",
				"%d"
			],
			"gasLimit": "200,000,000",
			"gasPrice": "0"
		}
	}`,
		user,
		swapPair.address,
		swapPair.firstToken,
		amountA,
		swapPair.secondToken,
		amountB,
	))

	if err == nil {
		statistics.swapFixedOutputHits += 1

		tokenAAfter, err := pfe.getTokens(user, swapPair.firstToken)
		if err != nil {
			return err
		}
		tokenBAfter, err := pfe.getTokens(user, swapPair.secondToken)
		if err != nil {
			return err
		}

		if amountInErr == nil {
			if tokenAAfter.Cmp(big.NewInt(0).Sub(tokenABefore, big.NewInt(0).SetBytes(amountInRaw[0]))) != 0 {
				return errors.New("swap fixed output wrong amount A")
			}
			if tokenBAfter.Cmp(big.NewInt(0).Add(tokenBBefore, big.NewInt(int64(amountB)))) != 0 {
				return errors.New("swap fixed output wrong amount B")
			}
		}
	} else {
		statistics.swapFixedOutputMisses += 1

		if output == nil {
			return errors.New("output is nil")
		}

		pfe.log("swapFixedOutput %s -> %s", swapPair.firstToken, swapPair.secondToken)
		pfe.log("could not swap because %s", output.ReturnMessage)

		expectedErrors := map[string]bool{
			"Insufficient reserve for token out": true,
			"Computed amount out lesser than minimum amount out": true,
			"Insufficient amount out reserve": true,
			"Optimal value is zero": true,
		}

		_, expected := expectedErrors[output.ReturnMessage]
		if !expected {
			return errors.New(output.ReturnMessage)
		}
	}

	return nil
}
