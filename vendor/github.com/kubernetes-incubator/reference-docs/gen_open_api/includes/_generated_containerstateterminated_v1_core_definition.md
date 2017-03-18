## ContainerStateTerminated v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | ContainerStateTerminated



ContainerStateTerminated is a terminated state of a container.

<aside class="notice">
Appears In  <a href="#containerstate-v1">ContainerState</a> </aside>

Field        | Description
------------ | -----------
containerID <br /> *string*  | Container's ID in the format 'docker://<container_id>'
exitCode <br /> *integer*  | Exit status from the last termination of the container
finishedAt <br /> *[Time](#time-v1)*  | Time at which the container last terminated
message <br /> *string*  | Message regarding the last termination of the container
reason <br /> *string*  | (brief) reason from the last termination of the container
signal <br /> *integer*  | Signal from the last termination of the container
startedAt <br /> *[Time](#time-v1)*  | Time at which previous execution of the container started

