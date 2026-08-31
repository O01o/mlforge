<script lang="ts">
    import uPlot from "uplot";
    import { onMount } from "svelte";
    import type { UPlotProps } from "$lib/types/uplot/uplot";
    import MLforgeSideBar from "$lib/components/common/MLforgeSideBar.svelte";
    import MLforgeSideBarItem from "$lib/components/common/MLforgeSideBarItem.svelte";
    import RunSidebarItem from "$lib/components/run/RunSidebarItem.svelte";
    import RunSummaryCard from "$lib/components/run/RunSummaryCard.svelte";
    import type { RunSummary, RunDetails } from "$lib/types/base/run";
    import type { RunDetailsActivation } from "$lib/types/uplot/run";
    import type { ExperimentSummary } from "$lib/types/base/experiment";
    import type { PlotRange } from "$lib/types/base/plot";
    import type { RunMetricPlots } from "$lib/types/base/plot";
    import { fetchExperimentDetails } from "$lib/api/experiment";
    import { page } from '$app/state';

    const experimentIdString = $derived(page.params.experiment ?? "-");
    const experimentId = Number(experimentIdString);
    if (isNaN(experimentId)) {
        throw new Error("Experiment ID is not a number");
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
    let chartContainer: HTMLDivElement | undefined = $state();

    let props: UPlotProps = $state({
        data: [
            [0, 1, 2, 3, 4, 5],
            [0, 1, 4, 9, 16, 25],
            [0, 1, 0.5, null, 0.125, 0.0625]
        ],
        options: {
            series: [
                {},
                {
                    label: "y = x^2",
                    stroke: "blue",
                    width: 2
                },
                {
                    label: "y = 1/x",
                    stroke: "green",
                    width: 2
                }
            ],
            title: "Sample Chart",
            width: 400,
            height: 300
        }
    });

    onMount(async () => {
        let experimentDetails = (await fetchExperimentDetails(experimentId)).experimentDetails;
        experimentSummary = experimentDetails.experimentSummary;
        for (let run of experimentDetails.runDetails) {
            runDetailActivations.push({
                runDetails: run,
                active: false
            });
        }
        new uPlot(props.options, props.data, chartContainer);
    });
</script>

<div class="d-flex vh-100 overflow-hidden">

    <MLforgeSideBar>
        <div>
            <button class="btn btn-primary">Get Result</button>
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
                    value={plotRange.startStep}
                />
            </div>
            <div>
                <input 
                    type="number" 
                    class="form-control" 
                    id="plotRangeEndFormControlInput" 
                    placeholder="Plot Range End"
                    min="1"
                    value={plotRange.endStep}
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

            <div class="row g-3">
                {#each runDetailActivations as runActivation}
                    <RunSummaryCard run={runActivation.runDetails} /> 
                {/each}
            </div>
        </div>
    </main>
</div>
