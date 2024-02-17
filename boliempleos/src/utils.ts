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
