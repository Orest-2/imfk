import { Container } from "@cloudflare/containers";

export class ImfkContainer extends Container {
	defaultPort = 1447;
	sleepAfter = "2m";

	override onStart() {
		console.log("ImfkContainer started");
	}

	override onStop() {
		console.log("ImfkContainer stopped");
	}

	override onError(error: unknown) {
		console.log(`ImfkContainer error: ${error}`);
	}
}

export default {
	async fetch(request, env, ctx): Promise<Response> {
		const id = env.IMFK_CONTAINER.idFromName("imfk");
		const container = env.IMFK_CONTAINER.get(id);

		const stub = container as unknown as ImfkContainer;
		await stub.startAndWaitForPorts();

		return stub.fetch(request);
	},
} satisfies ExportedHandler<Env>;
