<script lang="ts">
    import { onMount } from "svelte";
    import MLforgeSideBar from "$lib/components/common/MLforgeSideBar.svelte";
    import RunSidebarItem from "$lib/components/run/RunSidebarItem.svelte";
    import RunSummaryCard from "$lib/components/run/RunSummaryCard.svelte";
    import PlotsCard from "$lib/components/run/PlotsCard.svelte";
    import type { ExperimentSummary } from "$lib/types/base/experiment";
    import type { PlotRange, RunMetricPlots } from "$lib/types/base/plot";
    import type { RunDetailsActivation } from "$lib/types/uplot/run";
    import type { UPlotProps } from "$lib/types/uplot/uplot";
    import { getUPlotProps } from "$lib/types/uplot/uplot";
    import { fetchExperimentDetails } from "$lib/api/experiment";
    import { fetchPlots } from "$lib/api/plots";
    import { page } from "$app/state";

    function getExperimentId(): number {
        const experimentId = Number(page.params.experiment ?? "-");
        if (Number.isNaN(experimentId)) {
            throw new Error("Experiment ID is not a number");
        }
        return experimentId;
    }

    let experimentSummary: ExperimentSummary = $state({
        id: 0,
        name: "",
        description: "",
        createdAt: new Date(),
        updatedAt: new Date(),
    });
    let runDetailActivations: RunDetailsActivation[] = $state([]);
    let plotRange: PlotRange = $state({
        startStep: 0,
        endStep: 5,
    });
    let runMetricPlots: RunMetricPlots[] = $state([]);
    let props: UPlotProps[] = $state([]);

    async function getResult(): Promise<void> {
        const runIds = runDetailActivations
            .filter((runActivation) => runActivation.active)
            .map((runActivation) => runActivation.runDetails.runSummary.id);

        if (runIds.length === 0) {
            runMetricPlots = [];
            props = [];
            return;
        }

        const response = await fetchPlots({
            plotRange,
            runIds,
        });
        runMetricPlots = response.runMetricPlots;
        props = getUPlotProps(plotRange, runMetricPlots);
    }

    onMount(async () => {
        const experimentDetails = (await fetchExperimentDetails(getExperimentId())).experimentDetails;
        experimentSummary = experimentDetails.experimentSummary;
        runDetailActivations = experimentDetails.runDetails.map((runDetails) => ({
            runDetails,
            active: false,
        }));
    });
</script>

<div class="d-flex vh-100 overflow-hidden">
    <MLforgeSideBar>
        <div>
            <button
                class="btn btn-primary"
                onclick={getResult}
            >
                Get Result
            </button>
        </div>
        <hr />
        <div>
            <div class="mb-3">
                <h5 class="text-white">Plot Range</h5>
            </div>
            <div class="mb-3">
                <input
                    type="number"
                    class="form-control"
                    id="plotRangeStartFormControlInput"
                    placeholder="Plot Range Start"
                    min="0"
                    bind:value={plotRange.startStep}
                />
            </div>
            <div>
                <input
                    type="number"
                    class="form-control"
                    id="plotRangeEndFormControlInput"
                    placeholder="Plot Range End"
                    min="1"
                    bind:value={plotRange.endStep}
                />
            </div>
        </div>
        <hr />
        {#each runDetailActivations as runActivation}
            <RunSidebarItem run={runActivation.runDetails} bind:active={runActivation.active} />
        {/each}
    </MLforgeSideBar>

    <main class="flex-grow-1 overflow-auto p-4">
        <div class="container-fluid">
            <div class="d-flex justify-content-between align-items-center mb-4">
                <div>
                    <h1 class="h3 mb-1">{experimentSummary.name}</h1>
                    <p class="text-body-secondary mb-0">
                        Training runs for this experiment
                    </p>
                </div>

                <button class="btn btn-primary">
                    New RunDetails
                </button>
            </div>

            <div class="row g-3 mb-4">
                {#each runDetailActivations as runActivation}
                    <RunSummaryCard run={runActivation.runDetails} />
                {/each}
            </div>

            {#if props.length > 0}
                <div class="row g-3">
                    {#each props as plot, index (`${plot.options.title ?? 'plot'}-${index}`)}
                        <PlotsCard {plot} />
                    {/each}
                </div>
            {/if}
        </div>
    </main>
</div>
