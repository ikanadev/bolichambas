import { Company, SCAN_DATE } from "./utils";

export default async function getData(): Promise<Company[]> {
	const resp = await fetch(`/jobs_${SCAN_DATE}.json`);
	return await resp.json();
}
