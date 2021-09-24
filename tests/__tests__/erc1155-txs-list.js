const base = require("./erc-txs-base");

function erc1155TxList(account, token, count, cursor) {
  return '{\n' +
      '  account(address:'+JSON.stringify(account)+') {\n' +
      '    address\n' +
      '    erc1155TxList(token:'+JSON.stringify(token)+', tokenId: "0x1", txType:"TRANSFER", count: '+count+', cursor: '+JSON.stringify(cursor)+') {\n' +
      '      edges {\n' +
      '        trx {\n' +
      '          trxHash\n' +
      '          trxIndex\n' +
      '          tokenAddress\n' +
      '          trxType\n' +
      '          sender\n' +
      '          recipient\n' +
      '          tokenId\n' +
      '          amount\n' +
      '          timeStamp\n' +
      '        }' +
      '        cursor\n' +
      '      }\n' +
      '      pageInfo {\n' +
      '        first\n' +
      '        last' +
      '        hasNext\n' +
      '        hasPrevious\n' +
      '      }' +
      '      totalCount\n' +
      '    }\n' +
      '  }\n' +
      '}';
}

test("ERC1155 txs pagination forward - my testing contract, Jirka's txs", async () => {
  await base.testTrxPagination(3, async (count, cursor) => {
    return (await base.graphqlrq(erc1155TxList("0xcCdE9E1e6499b8f29b4b29973fC8a57282D22Db8", "0x99914d1212260685c64f5d14adc2d5b8ee075a8e", count, cursor))).account.erc1155TxList;
  });
});

test("ERC1155 txs pagination backward - my testing contract, Jirka's txs", async () => {
  await base.testTrxPagination(-3, async (count, cursor) => {
    return (await base.graphqlrq(erc1155TxList("0xcCdE9E1e6499b8f29b4b29973fC8a57282D22Db8", "0x99914d1212260685c64f5d14adc2d5b8ee075a8e", count, cursor))).account.erc1155TxList;
  });
});

test("ERC1155 txs pagination forward - big contract", async () => {
  await base.testTrxPagination(50, async (count, cursor) => {
    return (await base.graphqlrq(erc1155TxList("0x89916177Ce2f24bdc461a9c093295552e58b9095", "0x226DCAdE73DB5749a8e5CD868aa33461AB6b35a6", count, cursor))).account.erc1155TxList;
  });
});

test("ERC1155 txs pagination backward - big contract", async () => {
  await base.testTrxPagination(-50, async (count, cursor) => {
    return (await base.graphqlrq(erc1155TxList("0x89916177Ce2f24bdc461a9c093295552e58b9095", "0x226DCAdE73DB5749a8e5CD868aa33461AB6b35a6", count, cursor))).account.erc1155TxList;
  });
});
