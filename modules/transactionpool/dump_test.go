package transactionpool

import (
	"testing"
)

// testTransactionDumping calls TransactionSet and puts the transactions in a
// block that gets submitted to the state. If there is an error, the
// transaction set is known to be invalid.
func (tpt *tpoolTester) testTransactionDumping() {
	// Get the transaction set.
	tset, err := tpt.transactionPool.TransactionSet()
	if err != nil {
		tpt.assistant.Tester.Error(err)
	}

	// Add the transaction set to a block and check that it is valid in the
	// state by adding it to the state.
	b, err := tpt.assistant.MineCurrentBlock(tset)
	if err != nil {
		tpt.assistant.Tester.Error(err)
	}
	err = tpt.assistant.State.AcceptBlock(b)
	if err != nil {
		tpt.assistant.Tester.Error(err)
	}
}

// testSiacoinTransactionDump adds a handful of siacoin transactions to the
// transaction pool and then runs testTransactionDumping to see that the pool
// set follows the rules of the blockchain.
func (tpt *tpoolTester) testSiacoinTransactionDump() {
	tpt.addDependentSiacoinTransactionToPool()
	tpt.testTransactionDumping()
}

// TestSiacoinTransactionDump creates a tpoolTester and uses it to call
// testSiacoinTransactionDump.
func TestSiacoinTransactionDump(t *testing.T) {
	tpt := CreateTpoolTester(t)
	tpt.testSiacoinTransactionDump()
}
