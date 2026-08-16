<script lang="ts">
    import MLforgeSideBar from "$lib/components/common/MLforgeSideBar.svelte";
    import MLforgeSideBarItem from "$lib/components/common/MLforgeSideBarItem.svelte";
    import RunSidebarItem from "$lib/components/run/RunSidebarItem.svelte";
    import RunSummaryCard from "$lib/components/run/RunSummaryCard.svelte";
    import { type Run } from "$lib/entity/api";

    const runs = Array.from({ length: 8 }, (_, i) => ({
        id: i + 1,
        name: `Run ${i + 1}`,
        description: "Sample MLforge training run",
        failureReason: "",
        status: 1,
        startedAt: new Date(),
        endedAt: new Date(),
        parameters: [
            {
                id: 1,
                name: "Learning Rate",
                value: "0.001"
            },
            {
                id: 2,
                name: "Batch Size",
                value: "32"
            }
        ],
        metrics: [
            {
                id: 1,
                name: "Accuracy"
            },
            {
                id: 2,
                name: "Loss"
            }
        ]
    } as Run));
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
                />
            </div>
            <div>
                <input 
                    type="number" 
                    class="form-control" 
                    id="plotRangeEndFormControlInput" 
                    placeholder="Plot Range End"
                    min="1"
                />
            </div>
        </div>
        <hr />
        {#each runs as run}
            <RunSidebarItem run={run} active={run.id === 1} />
        {/each}
    </MLforgeSideBar>


    <main class="flex-grow-1 overflow-auto p-4">
        <div class="container-fluid">
            <div class="d-flex justify-content-between align-items-center mb-4">
                <div>
                    <h1 class="h3 mb-1">Experiment 1</h1>
                    <p class="text-body-secondary mb-0">
                        Training runs for this experiment
                    </p>
                </div>

                <button class="btn btn-primary">
                    New Run
                </button>
            </div>

            <div class="row g-3">
                {#each runs as run}
                    <RunSummaryCard {run} /> 
                {/each}
            </div>
        </div>
    </main>
</div>
