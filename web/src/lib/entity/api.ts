export type ExperimentItem = {
    id: string;
    name: string;
    description: string;
    createdAt: Date;
    updatedAt: Date;
};

export type RunItem = {
    id: string;
    name: string;
    description: string;
    createdAt: Date;
    updatedAt: Date;
};

export type ParameterItem = {
    id: string;
    name: string;
    value: string;
    createdAt: Date;
    updatedAt: Date;
};

export type MetricItem = {
    id: string;
    name: string;
    createdAt: Date;
    updatedAt: Date;
};

export type PlotItem = {
    step: number;
    value: number;
    createdAt: Date;
    updatedAt: Date;
};