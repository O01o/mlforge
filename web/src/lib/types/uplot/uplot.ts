import type { Options, AlignedData } from "uplot";
import type { RunDetails } from "$lib/types/base/run";
import type { RunMetricPlots } from "../base/plot";
import type { PlotRange } from "$lib/types/base/plot";

export type UPlotProps = {
    options: Options;
    data: AlignedData;
};

function getColor(id: number): string {
    const colors: string[] = [
        "gray",
        "maroon",
        "red",
        "purple",
        "fuchsia",
        "green",
        "lime",
        "olive",
        "navy",
        "blue",
        "teal",
        "aqua"
    ];
    return colors[id % colors.length];
}

export function getUPlotProps(
    plotRange: PlotRange,
    runDetails: RunDetails[], 
    plots: RunMetricPlots[],
): UPlotProps[] {
    let props: UPlotProps[] = [];
    let x_data: number[] = Array.from(
        { length: plotRange.endStep - plotRange.startStep + 1 }, 
        (_, i) => i + plotRange.startStep
    );
    for (let r of runDetails) {
        for (let m of r.metrics) {
            let options: Options = {
                series: [],
                title: m.name,
                width: 400,
                height: 300,
            };
            let y_datas: (number|null)[][] = [];
            for (let pl of plots) {
                if (pl.runId === r.runSummary.id) {
                    options.series = options.series || [];
                    options.series.push({
                        label: `${pl.runId}`,
                        stroke: getColor(pl.runId),
                        width: 2,
                    });
                    let y_data: (number|null)[] = Array(x_data.length).fill(null);
                    for (let mp of pl.metricPlots) {
                        if (mp.metricId === m.id) {
                            for (let sv of mp.plots) {
                                let index = sv[0] - plotRange.startStep;
                                if (index >= 0 && index < y_data.length) {
                                    y_data[index] = sv[1];
                                }
                            }
                            y_datas.push(y_data);
                        }
                    }
                    props.push({ options, data: [x_data, ...y_datas] });
                }
            }
        }
    }
    return props;
}