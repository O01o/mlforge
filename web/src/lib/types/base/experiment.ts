import type { RunDetails } from './run';

export type ExperimentSummary = {
    id: number;
    name: string;
    description: string;
    createdAt: Date;
    updatedAt: Date;
};

export type ExperimentDetails = {
    experimentSummary: ExperimentSummary;
    runDetails: RunDetails[];
};
