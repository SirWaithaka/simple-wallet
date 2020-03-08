package transaction

//For ministatement you need another table.. That will store transactions..
//It should have things like transaction ID, transaction type, timestamp,
//amount, user ID, etc.. Every time user deposits or withdraws you update
//that table accordingly.. So when a user requests for ministatement you
//give them the most recent 5 transactions.
