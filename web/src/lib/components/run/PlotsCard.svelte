<script lang="ts">
    import uPlot from "uplot";
    import type { UPlotProps } from "$lib/types/uplot/uplot";

    let {
        plot,
    }: {
        plot: UPlotProps;
    } = $props();

    let chartContainer: HTMLDivElement | undefined = $state();

    $effect(() => {
        if (!chartContainer) {
            return;
        }

        const chart = new uPlot(plot.options, plot.data, chartContainer);
        return () => {
            chart.destroy();
        };
    });
</script>

<div class="col-12 col-xl-6">
    <div class="card h-100">
        <div class="card-body">
            <div bind:this={chartContainer}></div>
        </div>
    </div>
</div>
