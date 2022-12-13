package dev.openfunction.samples.liberty.health;

import java.util.logging.Level;
import java.util.logging.Logger;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.core.Response;

@Path("/health")
public class HealthCheckResource {

	private static final String CLASS_NAME = HealthCheckResource.class.getName();
	private static final Logger LOGGER = Logger.getLogger(CLASS_NAME);

	private final HealthCheckService healthCheckService;

	public HealthCheckResource(HealthCheckService healthCheckService) {
		this.healthCheckService = healthCheckService;
	}

	@GET
	public Response check() {
		final String METHOD_NAME = "check";
		if (LOGGER.isLoggable(Level.FINER)) {
			LOGGER.entering(CLASS_NAME, METHOD_NAME);
		}

		final Response response = this.healthCheckService.isJvmHealthy() ? Response.ok().build()
				: Response.serverError().build();

		if (LOGGER.isLoggable(Level.FINER)) {
			LOGGER.exiting(CLASS_NAME, METHOD_NAME, response);
		}
		return response;
	}
}
