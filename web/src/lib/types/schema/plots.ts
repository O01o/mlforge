import type { RunMetricPlots, PlotRange } from '../base/plot';

export type RunMetricPlotsRequest = {
    plotRange: PlotRange;
    runIds: string[];
};

export type RunMetricPlotsResponse = {
    runMetricPlots: RunMetricPlots[];
};