import type { Options, AlignedData } from "uplot";
import type { RunMetricPlots, PlotRange } from "../base/plot";

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

function getPlotEntries(plots: Record<string, number>): [number, number][] {
    return Object.entries(plots).map(([step, value]) => [Number(step), value]);
}

export function getUPlotProps(
    plotRange: PlotRange,
    plots: RunMetricPlots[],
): UPlotProps[] {
    const props: UPlotProps[] = [];
    const xData: number[] = Array.from(
        { length: plotRange.endStep - plotRange.startStep + 1 },
        (_, i) => i + plotRange.startStep,
    );

    for (const p of plots) {
        const options: Options = {
            series: [{}],
            title: p.metricName,
            width: 400,
            height: 300,
        };
        const yDatas: (number | null)[][] = [];

        for (const m of p.metricPlots) {
            options.series.push({
                label: `${m.runId}`,
                stroke: getColor(m.runId),
                width: 2,
            });

            const yData: (number | null)[] = Array(xData.length).fill(null);
            for (const [step, value] of getPlotEntries(m.plots)) {
                const index = step - plotRange.startStep;
                if (index >= 0 && index < yData.length) {
                    yData[index] = value;
                }
            }
            yDatas.push(yData);
        }

        const data: AlignedData = [xData, ...yDatas];
        props.push({ options, data });
    }

    return props;
}
