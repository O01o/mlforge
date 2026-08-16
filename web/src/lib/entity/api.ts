export type Experiment = {
    id: number;
    name: string;
    description: string;
    createdAt: Date;
    updatedAt: Date;
};

export type Run = {
    id: number;
    name: string;
    description: string;
    failureReason: string;
    status: number;
    startedAt: Date;
    endedAt: Date;
    parameters: Parameter[];
    metrics: Metric[];
};

export type Parameter = {
    id: number;
    name: string;
    value: string;
};

export type Metric = {
    id: number;
    name: string;
};

export type Plot = {
    step: number;
    value: number;
};