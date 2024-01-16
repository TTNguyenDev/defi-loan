start: 
	ignite chain serve
loans:
	loand q loan list-loan
request:
	loand tx loan request-loan 1000token 100token 1000foocoin 500 --from alice --chain-id loan
approve: 
	loand tx loan approve-loan 0 --from bob --chain-id loan 
repay:
	loand tx loan repay-loan 0 --from alice --chain-id loan
liquidate:
	loand tx loan liquidate-loan 1 --from bob --chain-id loan
cancel:
	loand tx loan cancel-loan 0 --from alice --chain-id loan
