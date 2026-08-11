<script lang="ts">
    import MLforgeSideBar from "$lib/components/common/MLforgeSideBar.svelte";
    import MLforgeSideBarItem from "$lib/components/common/MLforgeSideBarItem.svelte";

    const runs = Array.from({ length: 30 }, (_, i) => ({
        id: i + 1,
        name: `Run ${i + 1}`,
        accuracy: (0.80 + i * 0.003).toFixed(3),
        loss: (0.50 - i * 0.01).toFixed(3)
    }));
</script>

<div class="d-flex vh-100 overflow-hidden">

    <MLforgeSideBar>
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
        <MLforgeSideBarItem path="/experiment1" active>
            Run 1
        </MLforgeSideBarItem>
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
                    <div class="col-12">
                        <div class="card">
                            <div class="card-body">
                                <div class="d-flex justify-content-between align-items-start">
                                    <div>
                                        <h2 class="h5 card-title">
                                            {run.name}
                                        </h2>

                                        <p class="card-text text-body-secondary">
                                            Sample MLforge training run
                                        </p>
                                    </div>

                                    <span class="badge text-bg-success">
                                        Finished
                                    </span>
                                </div>

                                <div class="row mt-3">
                                    <div class="col-md-4">
                                        <strong>Accuracy</strong>
                                        <div>{run.accuracy}</div>
                                    </div>

                                    <div class="col-md-4">
                                        <strong>Loss</strong>
                                        <div>{run.loss}</div>
                                    </div>

                                    <div class="col-md-4">
                                        <strong>Epochs</strong>
                                        <div>{run.id * 10}</div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                {/each}
            </div>
        </div>
    </main>
</div>
