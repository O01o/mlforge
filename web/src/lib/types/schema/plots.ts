import type { RunMetricPlots, PlotRange } from '../base/plot';

export type RunMetricPlotsRequest = {
    plotRange: PlotRange;
    runIds: number[];
};

export type RunMetricPlotsResponse = {
    runMetricPlots: RunMetricPlots[];
};
