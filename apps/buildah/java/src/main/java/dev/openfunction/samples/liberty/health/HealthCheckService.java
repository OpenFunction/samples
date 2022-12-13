package dev.openfunction.samples.liberty.health;

public interface HealthCheckService {

	/**
	 * Checks if the JVM is healthy
	 *
	 * @return true if it is healthy, otherwise false
	 */
	public boolean isJvmHealthy();

}