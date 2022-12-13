package dev.openfunction.samples.liberty;

import java.util.HashSet;
import java.util.Set;

import javax.ws.rs.ApplicationPath;
import javax.ws.rs.core.Application;

import dev.openfunction.samples.liberty.health.HealthCheckResource;
import dev.openfunction.samples.liberty.health.HealthCheckServiceImpl;

@ApplicationPath("/api")
public class LibertyApplication extends Application {

	@Override
	public Set<Object> getSingletons() {
		final Set<Object> set = new HashSet<>();
		set.add(new HealthCheckResource(new HealthCheckServiceImpl()));
		return set;
	}
}
