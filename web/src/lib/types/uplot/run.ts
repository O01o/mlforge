import type { RunDetails } from "$lib/types/base/run";

export type RunDetailsActivation = {
    active: boolean;
    runDetails: RunDetails;
}