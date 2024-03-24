import { ServiceDisruptor } from 'k6/x/disruptor';
import http from 'k6/http';

export function setup() {
    http.post(`http://localhost/bookstore/v1/purchase`, JSON.stringify([
        'Moby Dick; Or, The Whale',
        'Romeo and Juliet',
        'Frankenstein; Or, The Modern Prometheus',
    ]))
}

export default function (data) {
    http.get(`http://localhost/bookstore/v1/purchased`);
}

export function disrupt(data) {
    if (__ENV.SKIP_FAULTS == '1') {
        return;
    }

    const serviceDisruptor = new ServiceDisruptor('books-service', 'default');

    // delay traffic from the books service
    const fault = {
        averageDelay: '500ms',
        errorCode: 500,
        errorRate: 0.1,
    };
    serviceDisruptor.injectHTTPFaults(fault, '1m');
}

export const options = {
    scenarios: {
        load: {
            executor: 'constant-arrival-rate',
            rate: 100,
            preAllocatedVUs: 10,
            maxVUs: 100,
            exec: 'default',
            startTime: '0s',
            duration: '1m',
        },
        disrupt: {
            executor: 'shared-iterations',
            iterations: 1,
            vus: 1,
            exec: 'disrupt',
            startTime: '0s',
        },
    },
};