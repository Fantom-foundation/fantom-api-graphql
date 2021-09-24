const base = require("./erc-txs-base");

function erc721TxList(account, token, count, cursor) {
    return '{\n' +
        'erc721Transactions(token:'+JSON.stringify(token)+', txType:"TRANSFER", account:'+JSON.stringify(account)+', count: '+count+', cursor: '+JSON.stringify(cursor)+') {\n' +
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
        '}';
}

test("ERC721 txs pagination forward", async () => {
    await base.testTrxPagination(2, async (count, cursor) => {
        return (await base.graphqlrq(erc721TxList("0xb87097433709D7875e018c9Eec4fa4Ab1C0f800b", "0x7605cE039aFb683a2faC6877BdBf2Dbf36Da3dFe", count, cursor))).erc721Transactions;
    });
});

test("ERC721 txs pagination backward", async () => {
    await base.testTrxPagination(-2, async (count, cursor) => {
        return (await base.graphqlrq(erc721TxList("0xb87097433709D7875e018c9Eec4fa4Ab1C0f800b", "0x7605cE039aFb683a2faC6877BdBf2Dbf36Da3dFe", count, cursor))).erc721Transactions;
    });
});
