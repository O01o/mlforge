import type { ExperimentSummary, ExperimentDetails } from '../base/experiment';

export type ExperimentSummariesResponse = {
    experimentSummaries: ExperimentSummary[];
};

export type ExperimentDetailsResponse = {
    experimentDetails: ExperimentDetails;
};
