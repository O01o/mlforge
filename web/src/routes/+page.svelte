<script lang="ts">
    import Card from "$lib/components/common/Card.svelte";
    import uPlot from "uplot";
    import "uplot/dist/uPlot.min.css";
    import { onMount } from "svelte";
    import MLforgeSideBar from "$lib/components/common/MLforgeSideBar.svelte";
    import MLforgeSideBarItem from "$lib/components/common/MLforgeSideBarItem.svelte";
    import Title from "$lib/components/common/Title.svelte";
    import NamedCard from "$lib/components/common/NamedCard.svelte";
    import type { ExperimentSummary } from "$lib/types/base/experiment";
    import ExperimentSidebarItem from "$lib/components/experiment/ExperimentSidebarItem.svelte";
    import type { UPlotProps } from "$lib/types/uplot/plot";
    import type { AlignedData } from "uplot";
    
    const experiments = Array.from({ length: 30 }, (_, i) => ({
        id: i + 1,
        name: `ExperimentSummary ${i + 1}`,
        description: "Sample MLforge experiment",
    } as ExperimentSummary));

    
    let chartContainer: HTMLDivElement | undefined = $state();

    let props: UPlotProps = $state({
        data: [
            [0, 1, 2, 3, 4, 5],
            [0, 1, 4, 9, 16, 25]
        ],
        series: [
            {},
            {
                label: "y = x^2",
                stroke: "blue",
                width: 2
            }
        ],
        title: "Sample Chart",
        width: 800,
        height: 400
    });

    onMount(() => {
        const data: AlignedData = [
            [0, 1, 2, 3, 4, 5],
            [0, 1, 4, 9, 16, 25]
        ];

        const options = {
            title: "Sample Chart",
            width: 800,
            height: 400,
            series: [
                {},
                {
                    label: "y = x^2",
                    stroke: "blue",
                    width: 2
                }
            ]
        };

        new uPlot(options, data, chartContainer);
    });
</script>

<div class="d-flex vh-100 overflow-hidden">

    <MLforgeSideBar>
        {#each experiments as experiment}
            <ExperimentSidebarItem experiment={experiment} active={experiment.id === 1} />
        {/each}
    </MLforgeSideBar>

    <main class="flex-grow-1 overflow-auto p-4">
        <div class="container-fluid">
            <div class="d-flex justify-content-between align-items-center mb-4">
                <Title 
                    title="Create New ExperimentSummary" 
                    description="Create a new experiment to start tracking your training runs."
                />
                <button class="btn btn-primary">Create</button>
            </div>
            <div class="row g-3">
                <NamedCard
                    title="ExperimentSummary Name"
                    description="ExperimentSummary Name"
                >
                    <input 
                        type="text" 
                        class="form-control" 
                        id="experimentFormControlInput" 
                        placeholder="ExperimentSummary 1"
                    />
                </NamedCard>
                <NamedCard
                    title="ExperimentSummary Description"
                    description="ExperimentSummary Description"
                >
                    <textarea 
                        class="form-control" 
                        id="experimentFormControlTextarea" 
                        rows="3"
                    ></textarea>
                </NamedCard>

            </div>
        </div>
    </main>
</div>
