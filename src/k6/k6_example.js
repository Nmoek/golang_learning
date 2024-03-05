import http from "k6/http";

const url = 'http://localhost:8888/hello';

export default function () {
    const data = {name : 'ljk'}
    const cb200 = http.expectedStatuses(200)
    http.post(url, JSON.stringify(data), {
        headers: {'Content-type': 'application/json'},
        expectedStatus: cb200,
    });
}