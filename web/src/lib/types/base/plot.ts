export type Plot = {
    step: number;
    value: number;
};

export type PlotRange = {
    startStep: number;
    endStep: number;
};

export type MetricPlots = {
    metricId: number;
    plots: Map<number, number>;
};

export type RunMetricPlots = {
    runId: string;
    metricPlots: MetricPlots[];
};