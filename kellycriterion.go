package kellycriterion

//Calculate winChance should be less than 0; payout is the net odds received on the wager. (e.g. betting $10, on win, rewards $4 plus wager; then b=0.4)
func Calculate(winChance float64, payout float64) float64 {
	return (winChance*payout - (1 - winChance)) / payout
}
