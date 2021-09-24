const assert = require("assert");
const request = require("supertest");
const config = require("../config");

const api = request(config.config.API_SERVER);

exports.graphqlrq = async function(query) {
    const response = await api
        .post("/api")
        .set("Accept", "application/json")
        .send({ query: query })
        .expect(200);
    assert.strictEqual(response.body.errors, undefined, 'error from server!');
    return response.body.data;
}

exports.testTrxPagination = async function(count, queryFunc) {
    let hasNext, hasPrev;
    let listedCount = 0;
    let listedPages = 0;
    let totalCount = null;
    let cursor = null;
    let lastTimestamp = count < 0 ? 0 : Number.MAX_SAFE_INTEGER;
    let lastTimestampHex = null;

    do {
        let data = await queryFunc(count, cursor);
        if (totalCount == null) totalCount = data.totalCount;
        cursor = count > 0 ? data.pageInfo.last : data.pageInfo.first;
        hasNext = count > 0 ? data.pageInfo.hasNext : data.pageInfo.hasPrevious;
        hasPrev = count > 0 ? data.pageInfo.hasPrevious : data.pageInfo.hasNext;

        listedCount += data.edges.length;
        listedPages++;

        assert.strictEqual(hasPrev, listedPages !== 1, 'hasPrevious should be true for all except the first page');
        if (hasNext) {
            assert.strictEqual(data.edges.length, Math.abs(count), 'different amount of edges then requested count');
        }

        // check edges form an ordered list across pages - by timeStamp
        if (count > 0) {
            for (const edge of data.edges) {
                const timeStamp = parseInt(edge.trx.timeStamp);
                assert.ok(lastTimestamp >= timeStamp,
                    'incorrect order of edges! timeStamp '+edge.trx.timeStamp+' after '+lastTimestampHex+' page '+listedPages);
                lastTimestampHex = edge.trx.timeStamp;
                lastTimestamp = timeStamp;
            }
        } else {
            for (let i = data.edges.length-1; i >=0; i--) {
                const edge = data.edges[i];
                const timeStamp = parseInt(edge.trx.timeStamp);
                assert.ok(lastTimestamp <= timeStamp,
                    'incorrect order of edges! timeStamp '+edge.trx.timeStamp+' after '+lastTimestampHex+' page '+listedPages);
                lastTimestampHex = edge.trx.timeStamp;
                lastTimestamp = timeStamp;
            }
        }

    } while (hasNext);

    console.log('listedCount',listedCount, 'totalCount', parseInt(totalCount), 'pages', listedPages);
    assert.strictEqual(listedCount, parseInt(totalCount), 'listed different amount of txs then claim totalCount!');
    assert.ok(listedCount > Math.abs(count), 'not enough txs to test - '+listedCount);
}
