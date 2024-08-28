export type Job = {
	title: string;
	depto: string;
	url: string;
	publishDate: string;
	dueDate: string;
	content: string;
	area: string;
};

export type Company = {
	name: string;
	logoUrl: string;
	jobs: Job[];
}

export type JobItem = {
	company: string;
	companyLogo: string;
} & Job;

export const ALL_DEPTOS = "Todos";
export enum Depto {
	LaPaz = "La Paz",
	Oruro = "Oruro",
	Potosi = "Potosí",
	Cochabamba = "Cochabamba",
	Chuquisaca = "Chuquisaca",
	Tarija = "Tarija",
	SantaCruz = "Santa Cruz",
	Pando = "Pando",
	Beni = "Beni",
}

export type DeptoOption = Depto | typeof ALL_DEPTOS;

export const SCAN_DATE = "2024-08-28";
