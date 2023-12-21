
-- +migrate Up
INSERT INTO
	notification_flags ("group", "key", "description") 
VALUES 
	('wallet', 'disable_all', 'Disable all notification wallets'),
	('wallet', 'receive_transfer_success', 'Receive a tip'),
	('wallet', 'receive_airdrop_success', 'Receive airdrops'),
	('wallet', 'receive_deposit_success', 'Deposit completed'),
	('wallet', 'send_withdraw_success', 'Withdrawal completed'),
	('wallet', 'receive_payme_success', 'Payment request completed'),
	('wallet', '*_payme_expired', 'Payment request expired'),
	('wallet', '*_paylink_expired', 'Pay link has expired'),
	('wallet', 'send_paylink_success', 'Pay link claimed by another'),
	('wallet', 'receive_paylink_success', 'Claim a pay link'),
	('community', 'new_configuration', 'New configuration'),
	('app', 'new_vault_tx', 'New vault transactions'),
	('app', 'new_api_call', 'New API calls'),
	('app', 'info_updated', 'Information changes'),
	('app', 'new_member', 'New members');

-- +migrate Down
DELETE FROM notification_flags where key != '';
