const base = require("./erc-txs-base");

jest.setTimeout(60*1000);

function erc20TxList(account, token, count, cursor) {
  return '{\n' +
      '  account(address:'+JSON.stringify(account)+') {\n' +
      '    address\n' +
      '    erc20TxList(token: '+JSON.stringify(token)+', txType:"TRANSFER", count: '+count+', cursor: '+JSON.stringify(cursor)+') {\n' +
      '      edges {\n' +
      '        trx {\n' +
      '          trxHash\n' +
      '          trxIndex\n' +
      '          tokenAddress\n' +
      '          trxType\n' +
      '          sender\n' +
      '          recipient\n' +
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

test("ERC20 txs pagination forward", async () => {
  await base.testTrxPagination(50, async (count, cursor) => {
    return (await base.graphqlrq(erc20TxList("0x4824292dD7f9aCa724F267093cb340cCB629B5B1", "0x237Ed5598899e6fBdb1062D8556E28C3ECc400f8", count, cursor))).account.erc20TxList;
  });
});

test("ERC20 txs pagination backward", async () => {
  await base.testTrxPagination(-50, async (count, cursor) => {
    return (await base.graphqlrq(erc20TxList("0x4824292dD7f9aCa724F267093cb340cCB629B5B1", "0x237Ed5598899e6fBdb1062D8556E28C3ECc400f8", count, cursor))).account.erc20TxList;
  });
});
