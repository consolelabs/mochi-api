-- Language: sql
INSERT INTO offchain_tip_bot_tokens (id, token_id, token_name, token_symbol, icon, status, created_at, updated_at) VALUES 
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb1', '0x000', 'Ethereum', 'ETH', 'https://cdn.icon-icons.com/icons2/2107/PNG/512/file_type_ethereum_icon_130645.png', 1, now(), now()),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb2', '0x6b1', 'Dai', 'DAI', 'https://cdn.icon-icons.com/icons2/2107/PNG/512/file_type_dai_icon_130644.png', 1, now(), now()),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb3', '0x226', 'Wrapped Bitcoin', 'WBTC', 'https://cdn.icon-icons.com/icons2/2107/PNG/512/file_type_wbtc_icon_130646.png', 1, now(), now()),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb4', '0x0d81', 'Basic Attention Token', 'BAT', 'https://cdn.icon-icons.com/icons2/2107/PNG/512/file_type_bat_icon_130643.png', 1, now(), now()),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb5', '0x5142', 'Chainlink', 'LINK', 'https://cdn.icon-icons.com/icons2/2107/PNG/512/file_type_link_icon_130647.png', 1, now(), now()),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb6', '0x1f9', 'Uniswap', 'UNI', 'https://cdn.icon-icons.com/icons2/2107/PNG/512/file_type_uni_icon_130648.png', 1, now(), now()),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb7', '0x7fc', 'Aave', 'AAVE', 'https://cdn.icon-icons.com/icons2/2107/PNG/512/file_type_aave_icon_130642.png', 1, now(), now()),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb8', '0x1f5', 'Bancor', 'BNT', 'https://cdn.icon-icons.com/icons2/2107/PNG/512/file_type_bnt_icon_130645.png', 1, now(), now()),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb9', '0x9f8', 'Maker', 'MKR', 'https://cdn.icon-icons.com/icons2/2107/PNG/512/file_type_mkr_icon_130646.png', 1, now(), now()),
    ('6c4b78ab-48bf-47d4-b50f-55be43117b10', '0x408', 'Republic', 'REN', 'https://cdn.icon-icons.com/icons2/2107/PNG/512/file_type_ren_icon_130647.png', 1, now(), now()),
    ('6c4b78ab-48bf-47d4-b50f-55be43117b11', '0x0f5', 'Decentraland', 'MANA', 'https://cdn.icon-icons.com/icons2/2107/PNG/512/file_type_mana_icon_130646.png', 1, now(), now()),
    ('6c4b78ab-48bf-47d4-b50f-55be43117b12', '0x1ce', 'Compound', 'COMP', 'https://cdn.icon-icons.com/icons2/2107/PNG/512/file_type_comp_icon_130645.png', 1, now(), now()),
    ('6c4b78ab-48bf-47d4-b50f-55be43117b13', '0x514', 'Chainlink', 'LINK', 'https://cdn.icon-icons.com/icons2/2107/PNG/512/file_type_link_icon_130647.png', 1, now(), now()),
    ('6c4b78ab-48bf-47d4-b50f-55be43117b15', '0x0bc', 'Yearn Finance', 'YFI', 'https://cdn.icon-icons.com/icons2/2107/PNG/512/file_type_yfi_icon_130649.png', 1, now(), now());

INSERT INTO offchain_tip_bot_chains (id, chain_id, chain_name, currency, rpc_url, explorer_url, status, created_at, updated_at) VALUES 
    ('7303f2f8-b6d9-454d-aa92-880569fa5291', '0x14124', 'Ethereum', 'ETH', 'https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161', 'https://etherscan.io', 1, now(), now()),
    ('7303f2f8-b6d9-454d-aa92-880569fa5292', '0x342a4', 'Ropsten', 'ETH', 'https://ropsten.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161', 'https://ropsten.etherscan.io', 1, now(), now()),
    ('7303f2f8-b6d9-454d-aa92-880569fa5293', '0x38000', 'BSC', 'BNB', 'https://bsc-dataseed.binance.org/', 'https://bscscan.com', 1, now(), now()),
    ('7303f2f8-b6d9-454d-aa92-880569fa5294', '0x61000', 'BSC Testnet', 'BNB', 'https://data-seed-prebsc-1-s1.binance.org:8545/', 'https://testnet.bscscan.com', 1, now(), now()),
    ('7303f2f8-b6d9-454d-aa92-880569fa5295', '0x89000', 'Polygon', 'MATIC', 'https://rpc-mainnet.maticvigil.com/', 'https://polygonscan.com', 1, now(), now()),
    ('7303f2f8-b6d9-454d-aa92-880569fa5296', '0x81000', 'Fantom', 'FTM', 'https://rpcapi.fantom.network', 'https://ftmscan.com', 1, now(), now()),
    ('7303f2f8-b6d9-454d-aa92-880569fa5297', '0x4a000', 'xDai', 'xDAI', 'https://rpc.xdaichain.com/', 'https://blockscout.com/xdai/mainnet', 1, now(), now()),
    ('7303f2f8-b6d9-454d-aa92-880569fa5298', '0x64000', 'xDai Testnet', 'xDAI', 'https://rpc.xdaichain.com/', 'https://blockscout.com/xdai/mainnet', 1, now(), now()),
    ('7303f2f8-b6d9-454d-aa92-880569fa5299', '0x13881', 'Arbitrum', 'ARB', 'https://arb1.arbitrum.io/rpc', 'https://arbiscan.io', 1, now(), now()),
    ('7303f2f8-b6d9-454d-aa92-880569fa5210', '0x13880', 'Arbitrum Testnet', 'ARB', 'https://rinkeby.arbitrum.io/rpc', 'https://rinkeby-explorer.arbitrum.io', 1, now(), now());

INSERT INTO offchain_tip_bot_tokens_chains (token_id, chain_id) VALUES
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb1', '7303f2f8-b6d9-454d-aa92-880569fa5291'),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb2', '7303f2f8-b6d9-454d-aa92-880569fa5292'),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb3', '7303f2f8-b6d9-454d-aa92-880569fa5293'),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb4', '7303f2f8-b6d9-454d-aa92-880569fa5294'),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb5', '7303f2f8-b6d9-454d-aa92-880569fa5295'),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb6', '7303f2f8-b6d9-454d-aa92-880569fa5296'),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb7', '7303f2f8-b6d9-454d-aa92-880569fa5297'),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb8', '7303f2f8-b6d9-454d-aa92-880569fa5298'),
    ('6c4b78ab-48bf-47d4-b50f-55be43117bb9', '7303f2f8-b6d9-454d-aa92-880569fa5299'),
    ('6c4b78ab-48bf-47d4-b50f-55be43117b10', '7303f2f8-b6d9-454d-aa92-880569fa5210'),
    ('6c4b78ab-48bf-47d4-b50f-55be43117b11', '7303f2f8-b6d9-454d-aa92-880569fa5291'),
    ('6c4b78ab-48bf-47d4-b50f-55be43117b12', '7303f2f8-b6d9-454d-aa92-880569fa5292'),
    ('6c4b78ab-48bf-47d4-b50f-55be43117b13', '7303f2f8-b6d9-454d-aa92-880569fa5293'),
    ('6c4b78ab-48bf-47d4-b50f-55be43117b15', '7303f2f8-b6d9-454d-aa92-880569fa5294');

INSERT INTO offchain_tip_bot_contracts (id, chain_id, contract_address, status, assign_status, centralize_wallet, created_at, updated_at) VALUES
    ('4ab3c60d-4f2d-4c05-81f2-ba69860037e1', '7303f2f8-b6d9-454d-aa92-880569fa5291', '0x6c4b78ab48bf47d4b50f55be43117b11', 1, 1, '0x6c4b78ab48bf47d4b50f55be43117b10', now(), now()),
    ('4ab3c60d-4f2d-4c05-81f2-ba69860037e4', '7303f2f8-b6d9-454d-aa92-880569fa5292', '0x6c4b78ab48bf47d4b50f55be43117b14', 1, 1, '0x6c4b78ab48bf47d4b50f55be43117b11', now(), now()),
    ('4ab3c60d-4f2d-4c05-81f2-ba69860037e6', '7303f2f8-b6d9-454d-aa92-880569fa5293', '0x6c4b78ab48bf47d4b50f55be43117b16', 1, 1, '0x6c4b78ab48bf47d4b50f55be43117b12', now(), now()),
    ('4ab3c60d-4f2d-4c05-81f2-ba69860037e8', '7303f2f8-b6d9-454d-aa92-880569fa5294', '0x6c4b78ab48bf47d4b50f55be43117b18', 1, 1, '0x6c4b78ab48bf47d4b50f55be43117b13', now(), now()),
    ('4ab3c60d-4f2d-4c05-81f2-ba69860037e9', '7303f2f8-b6d9-454d-aa92-880569fa5295', '0x6c4b78ab48bf47d4b50f55be43117b19', 1, 1, '0x6c4b78ab48bf47d4b50f55be43117b14', now(), now()),
    ('4ab3c60d-4f2d-4c05-81f2-ba6986003711', '7303f2f8-b6d9-454d-aa92-880569fa5296', '0x6c4b78ab48bf47d4b50f55be43117b21', 1, 1, '0x6c4b78ab48bf47d4b50f55be43117b20', now(), now()),
    ('4ab3c60d-4f2d-4c05-81f2-ba6986003713', '7303f2f8-b6d9-454d-aa92-880569fa5297', '0x6c4b78ab48bf47d4b50f55be43117b33', 1, 1, '0x6c4b78ab48bf47d4b50f55be43117b21', now(), now()),
    ('4ab3c60d-4f2d-4c05-81f2-ba6986003715', '7303f2f8-b6d9-454d-aa92-880569fa5298', '0x6c4b78ab48bf47d4b50f55be43117b45', 1, 1, '0x6c4b78ab48bf47d4b50f55be43117b32', now(), now()),
    ('4ab3c60d-4f2d-4c05-81f2-ba6986003717', '7303f2f8-b6d9-454d-aa92-880569fa5299', '0x6c4b78ab48bf47d4b50f55be43117b57', 1, 1, '0x6c4b78ab48bf47d4b50f55be43117b43', now(), now()),
    ('4ab3c60d-4f2d-4c05-81f2-ba6986003719', '7303f2f8-b6d9-454d-aa92-880569fa5210', '0x6c4b78ab48bf47d4b50f55be43117b69', 1, 1, '0x6c4b78ab48bf47d4b50f55be43117b54', now(), now());