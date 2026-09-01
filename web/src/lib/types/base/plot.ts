export type Plot = {
    step: number;
    value: number;
};

export type PlotRange = {
    startStep: number;
    endStep: number;
};

export type MetricPlots = {
    runId: number;
    metricId: number;
    plots: Record<string, number>;
};

export type RunMetricPlots = {
    metricName: string;
    metricPlots: MetricPlots[];
};