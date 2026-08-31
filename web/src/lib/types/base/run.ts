import type { Parameter } from './parameter';
import type { Metric } from './metric';

export type RunSummary = {
    id: number;
    name: string;
    description: string;
    failureReason: string;
    status: number;
    startedAt?: Date;
    endedAt?: Date;
};

export type RunDetails = {
    runSummary: RunSummary;
    parameters: Parameter[];
    metrics: Metric[];
};