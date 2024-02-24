package org.example;

import org.gradle.api.Plugin;
import org.gradle.api.Project;

public class GreetingPlugin implements Plugin<Project> {

    @Override
    public void apply( Project project ) {
        System.out.println("Hello"); // test
    }
}
