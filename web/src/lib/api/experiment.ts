const fetchExperiments = async (): Promise<ExperimentSummary[]> => {
  const response = await fetch('/api/experiments');
  if (!response.ok) {
    throw new Error('Failed to fetch experiments');
  }
  return response.json();
};

export default fetchExperiments;