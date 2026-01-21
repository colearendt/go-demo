import { OpenAPIRoute } from "chanfana";
import { z } from "zod";

export class DataFetch extends OpenAPIRoute {
	schema = {
		tags: ["Data"],
		summary: "Get consumer goods price data",
		responses: {
			"200": {
				description: "Returns price data for consumer goods",
				content: {
					"application/json": {
						schema: z.object({
							labels: z.array(z.string()),
							datasets: z.array(
								z.object({
									label: z.string(),
									data: z.array(z.number()),
									borderColor: z.string(),
									backgroundColor: z.string(),
								})
							),
						}),
					},
				},
			},
		},
	};

	async handle(c: any) {
		return {
			labels: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
			datasets: [
				{
					label: "Food & Beverages",
					data: [100, 103, 107, 111, 114, 117, 120, 123, 125, 127, 129, 132],
					borderColor: "#667eea",
					backgroundColor: "rgba(102, 126, 234, 0.1)",
				},
				{
					label: "Personal Care",
					data: [100, 102, 104, 106, 108, 110, 112, 114, 116, 117, 118, 120],
					borderColor: "#764ba2",
					backgroundColor: "rgba(118, 75, 162, 0.1)",
				},
				{
					label: "Household Items",
					data: [100, 102, 105, 107, 109, 111, 113, 115, 116, 117, 118, 119],
					borderColor: "#f093fb",
					backgroundColor: "rgba(240, 147, 251, 0.1)",
				},
				{
					label: "Electronics",
					data: [100, 98, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87],
					borderColor: "#4facfe",
					backgroundColor: "rgba(79, 172, 254, 0.1)",
				},
				{
					label: "Clothing",
					data: [100, 101, 103, 105, 107, 109, 111, 113, 114, 115, 116, 117],
					borderColor: "#43e97b",
					backgroundColor: "rgba(67, 233, 123, 0.1)",
				},
				{
					label: "Transportation",
					data: [100, 105, 109, 113, 117, 120, 123, 125, 127, 129, 131, 134],
					borderColor: "#fa709a",
					backgroundColor: "rgba(250, 112, 154, 0.1)",
				},
			],
		};
	}
}
