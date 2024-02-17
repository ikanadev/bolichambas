import { Company } from "./utils";

export default async function getData(): Promise<Company[]> {
	const resp = await fetch("/jobs_2024-02-17.json");
	return await resp.json();
}
