import type { 
    RunMetricPlotsRequest,
    RunMetricPlotsResponse
} from '$lib/types/schema/plots';

export const fetchPlots = async (request: RunMetricPlotsRequest): Promise<RunMetricPlotsResponse> => {
  const response = await fetch(
    '/api/plots', 
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(request)
    }
  );
  if (!response.ok) {
    throw new Error('Failed to fetch plots');
  }
  return response.json();
};