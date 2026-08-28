import type { 
  ExperimentSummariesResponse, 
  ExperimentDetailsResponse 
} from '$lib/types/schema/experiment';

export const fetchExperimentSummaries = async (): Promise<ExperimentSummariesResponse> => {
  const response = await fetch('/api/experiments');
  if (!response.ok) {
    throw new Error('Failed to fetch experiments');
  }
  return response.json();
};

export const fetchExperimentDetails = async (experimentId: number): Promise<ExperimentDetailsResponse> => {
  const response = await fetch(`/api/experiments/${experimentId}`);
  if (!response.ok) {
    throw new Error(`Failed to fetch details for experiment ${experimentId}`);
  }
  return response.json();
};
