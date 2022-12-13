package dev.openfunction.samples.liberty.health;

import java.lang.management.ManagementFactory;
import java.lang.management.MemoryMXBean;
import java.lang.management.MemoryUsage;
import java.util.logging.Level;
import java.util.logging.Logger;

public class HealthCheckServiceImpl implements HealthCheckService {

	private static final String CLASS_NAME = HealthCheckServiceImpl.class.getName();
	private static final Logger LOGGER = Logger.getLogger(CLASS_NAME);
	private static final MemoryMXBean MEMORY_BEAN = ManagementFactory.getMemoryMXBean();

	@Override
	public boolean isJvmHealthy() {
		final String METHOD_NAME = "isJvmHealthy";
		if (LOGGER.isLoggable(Level.FINER)) {
			LOGGER.entering(CLASS_NAME, METHOD_NAME);
		}

		// retrieve the heap memory usage
		final MemoryUsage memoryUsage = MEMORY_BEAN.getHeapMemoryUsage();
		final long memUsed = memoryUsage.getUsed();
		final long memMax = memoryUsage.getMax();

		// assume it is healthy if at most 90 % of the heap are used
		final boolean healthy = memUsed < memMax * 0.9;

		if (LOGGER.isLoggable(Level.FINER)) {
			LOGGER.exiting(CLASS_NAME, METHOD_NAME, healthy);
		}
		return healthy;
	}
}
