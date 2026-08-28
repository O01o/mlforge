export type UPlotProps = {
    data: (number | null)[][] | (Float64Array | Float32Array)[];
    series: any[];
    title?: string;
    width?: number;
    height?: number;
};