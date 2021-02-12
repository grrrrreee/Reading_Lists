## How to scan the latest pairs added to Sushiswap ?

출처 : https://medium.com/coinmonks/how-to-scan-the-latest-pairs-added-to-sushiswap-4e5fee7880e2

Update: for those interested on having the latest pairs for Uniswap -> go Here

### What is Sushiswap ?

SushiSwap is a software running on Ethereum that seeks to incentivize a network of users to operate a platform where users can buy and sell crypto assets. Similar to platforms like Uniswap and Balancer, SushiSwap uses a collection of liquidity pools to achieve this goal. Users first lock up assets into smart contracts, and traders then buy and sell cryptocurrencies from those pools, swapping out one token for another.

One of a growing number of decentralized finance (DeFi) platforms, SushiSwap allows users to trade cryptocurrencies without the need for a central operator administrator.

This means that decisions relating to the SushiSwap software are made by holders of its native cryptocurrency, SUSHI. Anyone holding a balance of the asset can propose changes to how it operates, and can vote on submitted proposals by other users.

image

It’s core function is to mirror a traditional exchange by facilitating the buying and selling of different crypto assets between users. Rather than being supported by one central entity, tokens traded on SushiSwap are maintained by smart contracts, and users lock crypto on the software that can then be accessed by traders.

Rather than being supported by one central entity, tokens traded on SushiSwap are maintained by smart contracts, and users lock crypto on the software that can then be accessed by traders.

Of note, those who trade against locked assets pay a fee that is then distributed to all liquidity providers proportionally, based on their contribution to the pool.

This is how Sushiswap’s Decentralized Exchange looks like:

image

Some of the CryptoTools data analytics users have been interested in finding a way to get the latest tokens trading on Sushiswap and Uniswap, most probably as a trading analytics tool.

As a result, I’ve created a Google Sheet templates that helps you filter new tradable coins on these 2 decentralized exchanges.

ACCESS LIVE TEMPLATE SHEET HERE

The sheet returns all new tradable pairs on Sushiswap and Uniswap, giving constraints on the Number of Days the pair has been active, the Volume ($), the Liquidity ($), and the number of Transactions.

image

In order to get Sushiswap’s analytics I used The Graph which is an indexing protocol for querying networks like Ethereum and IPFS. Anyone can use, build and publish open APIs, called subgraphs, making data easily accessible.

image

SUSHISWAP FUNCTION IN GOOGLE SHEETS:
Returns new tradable pairs on Sushiswap, giving constraints on the Number of Days the coin is active, the Volume ($), the Liquidity ($), and the number of Transactions .

image

For example, if I want to get the new Sushiswap pairs where:
- the pool was launched in the last 4 Days
- the daily Volume is greater than $5'000
-  the Liquidity is above $10'000
and there has been more than 200 Transactions since the launch

The formula becomes: =SUSHISWAP(4,5000,10000,200)

@param {days} the number of Days since the pair is active
@param {volume} the minimum Volume ($)
@param {liquidity} the minimum Liquidity ($)
@param {tx_count} the number of Transactions existant since creation

* @return a table (see GIF above)with all new tradable pairs on Uniswap and their number of Days since Active, the Volume ($), the Liquidity ($), the number of Transactions.

### More indicators for scanning ?

There are plenty more functionalities that can be added through the TheGraph API. Don’t hesitate to have a look at all available end points like:

image 

- totalSupply
- untrackedVolumeUSD
-  liquidityProviderCount

If you are interested in getting some help in integrating more personalized indicators, DM me.

### Conclusion

An easy accessible way to get the latest pairs on Sushiswap using Google Sheets which can be used for data analytics, and as a trading tools for screening of new market participants.