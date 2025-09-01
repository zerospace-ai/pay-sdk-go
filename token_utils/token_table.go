package token_utils

const tokenJson = `[
  {
    "mainnet": "Ethereum",
    "chain_id": 1,
    "token_id": 2,
    "token_address": "0xdac17f958d2ee523a2206206994597c13d831ec7",
    "symbol": "USDT-ERC20",
    "decimals": 6
  },
  {
    "mainnet": "Ethereum",
    "chain_id": 1,
    "token_id": 140,
    "token_address": "0x6982508145454Ce325dDbE47a25d4ec3d2311933",
    "symbol": "PEPE",
    "decimals": 18
  },
  {
    "mainnet": "Ethereum",
    "chain_id": 1,
    "token_id": 141,
    "token_address": "0xb131f4A55907B10d1F0A50d8ab8FA09EC342cd74",
    "symbol": "MEME",
    "decimals": 18
  },
  {
    "mainnet": "Ethereum",
    "chain_id": 1,
    "token_id": 64,
    "token_address": "0xEd04915c23f00A313a544955524EB7DBD823143d",
    "symbol": "ACH",
    "decimals": 8
  },
  {
    "mainnet": "Ethereum",
    "chain_id": 1,
    "token_id": 25,
    "token_address": "0x6B175474E89094C44Da98b954EedeAC495271d0F",
    "symbol": "DAI",
    "decimals": 18
  },
  {
    "mainnet": "Ethereum",
    "chain_id": 1,
    "token_id": 24,
    "token_address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
    "symbol": "USDC",
    "decimals": 6
  },
  {
    "mainnet": "Ethereum",
    "chain_id": 1,
    "token_id": 142,
    "token_address": "0x163f8C2467924be0ae7B5347228CABF260318753",
    "symbol": "WLD",
    "decimals": 18
  },
  {
    "mainnet": "Ethereum",
    "chain_id": 1,
    "token_id": 1,
    "token_address": "ETH",
    "symbol": "ETH",
    "decimals": 18
  },
  {
    "mainnet": "Tron",
    "chain_id": 2,
    "token_id": 40,
    "token_address": "THb4CqiFdwNHsWsQCs4JhzwjMWys4aqCbF",
    "symbol": "ETH",
    "decimals": 18
  },
  {
    "mainnet": "Tron",
    "chain_id": 2,
    "token_id": 90,
    "token_address": "TPYmHEhy5n8TCEfYGqW2rPxsghSfzghPDn",
    "symbol": "USDD",
    "decimals": 18
  },
  {
    "mainnet": "Tron",
    "chain_id": 2,
    "token_id": 26,
    "token_address": "TEkxiTehnzSmSe2XqrBj4w32RUN966rdz8",
    "symbol": "USDC",
    "decimals": 6
  },
  {
    "mainnet": "Tron",
    "chain_id": 2,
    "token_id": 33,
    "token_address": "TSkW873XMKiDCxGZrA4YH8KGeipLdC6Gyu",
    "symbol": "CVNT",
    "decimals": 18
  },
  {
    "mainnet": "Tron",
    "chain_id": 2,
    "token_id": 3,
    "token_address": "TRX",
    "symbol": "TRX",
    "decimals": 6
  },
  {
    "mainnet": "Tron",
    "chain_id": 2,
    "token_id": 4,
    "token_address": "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
    "symbol": "USDT-TRC20",
    "decimals": 6
  },
  {
    "mainnet": "Bitcoin",
    "chain_id": 3,
    "token_id": 100,
    "token_address": "BTC",
    "symbol": "BTC",
    "decimals": 8
  },
  {
    "mainnet": "Bitcoin",
    "chain_id": 3,
    "token_id": 102,
    "token_address": "SATS",
    "symbol": "SATS",
    "decimals": 18
  },
  {
    "mainnet": "Bitcoin",
    "chain_id": 3,
    "token_id": 103,
    "token_address": "RATS",
    "symbol": "RATS",
    "decimals": 18
  },
  {
    "mainnet": "Bitcoin",
    "chain_id": 3,
    "token_id": 101,
    "token_address": "ORDI",
    "symbol": "ORDI",
    "decimals": 18
  },
  {
    "mainnet": "Solana",
    "chain_id": 4,
    "token_id": 400,
    "token_address": "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB",
    "symbol": "USDT",
    "decimals": 6
  },
  {
    "mainnet": "Solana",
    "chain_id": 4,
    "token_id": 401,
    "token_address": "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
    "symbol": "USDC",
    "decimals": 6
  },
  {
    "mainnet": "Solana",
    "chain_id": 4,
    "token_id": 19,
    "token_address": "SOL",
    "symbol": "SOL",
    "decimals": 9
  },
  {
    "mainnet": "Solana",
    "chain_id": 4,
    "token_id": 410,
    "token_address": "nQ1qgSpXWi71twnWPFjyfCtcbUXbVyQb64RfHKwRpKE",
    "symbol": "DAOT",
    "decimals": 9
  },
  {
    "mainnet": "XRP",
    "chain_id": 5,
    "token_id": 200,
    "token_address": "XRP",
    "symbol": "XRP",
    "decimals": 6
  },
  {
    "mainnet": "DogeCoin",
    "chain_id": 9,
    "token_id": 300,
    "token_address": "DOGE",
    "symbol": "DOGE",
    "decimals": 8
  },
  {
    "mainnet": "Optimistic",
    "chain_id": 10,
    "token_id": 131,
    "token_address": "0xdC6fF44d5d932Cbd77B52E5612Ba0529DC6226F1",
    "symbol": "WLD",
    "decimals": 18
  },
  {
    "mainnet": "Optimistic",
    "chain_id": 10,
    "token_id": 130,
    "token_address": "ETH",
    "symbol": "ETH",
    "decimals": 18
  },
  {
    "mainnet": "Optimistic",
    "chain_id": 10,
    "token_id": 133,
    "token_address": "0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85",
    "symbol": "USDC",
    "decimals": 6
  },
  {
    "mainnet": "Optimistic",
    "chain_id": 10,
    "token_id": 132,
    "token_address": "0x94b008aA00579c1307B0EF2c499aD98a8ce58e58",
    "symbol": "USDT",
    "decimals": 6
  },
  {
    "mainnet": "Bnb Smart Chain",
    "chain_id": 56,
    "token_id": 62,
    "token_address": "0xc0be866ecc026957fc7160c1a45f2bee9870fd46",
    "symbol": "ARK",
    "decimals": 18
  },
  {
    "mainnet": "Bnb Smart Chain",
    "chain_id": 56,
    "token_id": 68,
    "token_address": "0x6FDcdfef7c496407cCb0cEC90f9C5Aaa1Cc8D888",
    "symbol": "VET",
    "decimals": 18
  },
  {
    "mainnet": "Bnb Smart Chain",
    "chain_id": 56,
    "token_id": 63,
    "token_address": "0x8540f3D726Aed340Bc57Fd07a61b0ae2a9d5ECa9",
    "symbol": "PUC",
    "decimals": 18
  },
  {
    "mainnet": "Bnb Smart Chain",
    "chain_id": 56,
    "token_id": 65,
    "token_address": "0xbc7d6b50616989655afd682fb42743507003056d",
    "symbol": "ACH",
    "decimals": 8
  },
  {
    "mainnet": "Bnb Smart Chain",
    "chain_id": 56,
    "token_id": 66,
    "token_address": "0xFE8bF5B8F5e4eb5f9BC2be16303f7dAB8CF56aA8",
    "symbol": "BIBI",
    "decimals": 18
  },
  {
    "mainnet": "Bnb Smart Chain",
    "chain_id": 56,
    "token_id": 29,
    "token_address": "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",
    "symbol": "BUSD",
    "decimals": 18
  },
  {
    "mainnet": "Bnb Smart Chain",
    "chain_id": 56,
    "token_id": 31,
    "token_address": "0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c",
    "symbol": "BTCB",
    "decimals": 18
  },
  {
    "mainnet": "Bnb Smart Chain",
    "chain_id": 56,
    "token_id": 30,
    "token_address": "0x2170Ed0880ac9A755fd29B2688956BD959F933F8",
    "symbol": "ETH",
    "decimals": 18
  },
  {
    "mainnet": "Bnb Smart Chain",
    "chain_id": 56,
    "token_id": 6,
    "token_address": "0x55d398326f99059ff775485246999027b3197955",
    "symbol": "USDT-BEP20",
    "decimals": 18
  },
  {
    "mainnet": "Bnb Smart Chain",
    "chain_id": 56,
    "token_id": 23,
    "token_address": "0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3",
    "symbol": "DAI",
    "decimals": 18
  },
  {
    "mainnet": "Bnb Smart Chain",
    "chain_id": 56,
    "token_id": 22,
    "token_address": "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d",
    "symbol": "USDC",
    "decimals": 18
  },
  {
    "mainnet": "Bnb Smart Chain",
    "chain_id": 56,
    "token_id": 5,
    "token_address": "BNB",
    "symbol": "BNB",
    "decimals": 18
  },
  {
    "mainnet": "Polygon",
    "chain_id": 137,
    "token_id": 12,
    "token_address": "0xc2132D05D31c914a87C6611C10748AEb04B58e8F",
    "symbol": "USDT",
    "decimals": 6
  },
  {
    "mainnet": "Polygon",
    "chain_id": 137,
    "token_id": 13,
    "token_address": "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174",
    "symbol": "USDC",
    "decimals": 6
  },
  {
    "mainnet": "Polygon",
    "chain_id": 137,
    "token_id": 11,
    "token_address": "MATIC",
    "symbol": "MATIC",
    "decimals": 18
  },
  {
    "mainnet": "Polygon",
    "chain_id": 137,
    "token_id": 110,
    "token_address": "0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359",
    "symbol": "USDC",
    "decimals": 6
  },
  {
    "mainnet": "CVN",
    "chain_id": 2032,
    "token_id": 7,
    "token_address": "CVN",
    "symbol": "CVN",
    "decimals": 18
  },
  {
    "mainnet": "CVN",
    "chain_id": 2032,
    "token_id": 35,
    "token_address": "0x109B57A29eE6E9A93f33687F6CE553fB18D8EE78",
    "symbol": "USDT-CRC20",
    "decimals": 6
  },
  {
    "mainnet": "CVN",
    "chain_id": 2032,
    "token_id": 51,
    "token_address": "0x6b94b0a2878c68811c1bd6cecc2b7cc44a9ed7ab",
    "symbol": "HPT",
    "decimals": 8
  },
  {
    "mainnet": "Merlin",
    "chain_id": 4200,
    "token_id": 500,
    "token_address": "BTC",
    "symbol": "BTC",
    "decimals": 18
  },
  {
    "mainnet": "Merlin",
    "chain_id": 4200,
    "token_id": 501,
    "token_address": "0x5c46bFF4B38dc1EAE09C5BAc65872a1D8bc87378",
    "symbol": "MERL",
    "decimals": 18
  },
  {
    "mainnet": "Base",
    "chain_id": 8453,
    "token_id": 801,
    "token_address": "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913",
    "symbol": "USDC",
    "decimals": 6
  },
  {
    "mainnet": "Base",
    "chain_id": 8453,
    "token_id": 802,
    "token_address": "ETH",
    "symbol": "ETH",
    "decimals": 18
  },
  {
    "mainnet": "TON",
    "chain_id": null,
    "token_id": 201,
    "token_address": "0:105e5589bc66db15f13c177a12f2cf3b94881da2f4b8e7922c58569176625eb5",
    "symbol": "JETTON",
    "decimals": 9
  },
  {
    "mainnet": "TON",
    "chain_id": 15186,
    "token_id": 202,
    "token_address": "0:b113a994b5024a16719f69139328eb759596c38a25f59028b146fecdc3621dfe",
    "symbol": "USDT",
    "decimals": 6
  },
  {
    "mainnet": "TON",
    "chain_id": 15186,
    "token_id": 200,
    "token_address": "TON",
    "symbol": "TON",
    "decimals": 9
  },
  {
    "mainnet": "Arbitrum One",
    "chain_id": 42161,
    "token_id": 122,
    "token_address": "0xaf88d065e77c8cC2239327C5EDb3A432268e5831",
    "symbol": "USDC",
    "decimals": 6
  },
  {
    "mainnet": "Arbitrum One",
    "chain_id": 42161,
    "token_id": 121,
    "token_address": "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9",
    "symbol": "USDT",
    "decimals": 6
  },
  {
    "mainnet": "Arbitrum One",
    "chain_id": 42161,
    "token_id": 120,
    "token_address": "ETH",
    "symbol": "ETH",
    "decimals": 18
  },
  {
    "mainnet": "Arbitrum One",
    "chain_id": 42161,
    "token_id": 123,
    "token_address": "0x9fE175843Df9deCd99C78E72b2424C47D61Ad2bF",
    "symbol": "ATM",
    "decimals": 18
  },
  {
    "mainnet": "Arbitrum One",
    "chain_id": 42161,
    "token_id": 124,
    "token_address": "0x58BDf739aE17d1C60C6FD3433E288E38B81C2853",
    "symbol": "SAM",
    "decimals": 18
  },
  {
    "mainnet": "Avax Chain C",
    "chain_id": 43114,
    "token_id": 18,
    "token_address": "0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E",
    "symbol": "USDC",
    "decimals": 6
  },
  {
    "mainnet": "Avax Chain C",
    "chain_id": 43114,
    "token_id": 17,
    "token_address": "0xc7198437980c041c805A1EDcbA50c1Ce5db95118",
    "symbol": "USDT",
    "decimals": 6
  },
  {
    "mainnet": "Avax Chain C",
    "chain_id": 43114,
    "token_id": 16,
    "token_address": "AVAX",
    "symbol": "AVAX",
    "decimals": 18
  },
  {
    "mainnet": "NA Chain",
    "chain_id": 65143,
    "token_id": 600,
    "token_address": "NAC",
    "symbol": "NAC",
    "decimals": 9
  },
  {
    "mainnet": "NA Chain",
    "chain_id": 65143,
    "token_id": 601,
    "token_address": "GAT",
    "symbol": "GAT",
    "decimals": 9
  },
  {
    "mainnet": "ODIN",
    "chain_id": 666666,
    "token_id": 80,
    "token_address": "ODIN",
    "symbol": "ODIN",
    "decimals": 18
  },
  {
    "mainnet": "THOR",
    "chain_id": 868868,
    "token_id": 81,
    "token_address": "THOR",
    "symbol": "THOR",
    "decimals": 18
  }
]
`
