# distributed-interrupts
When you need to interrupt a task on another instance

## why this ##
Imagine you have a tasks running on few instances but there are other instances that require resources already utilized to run its own task. <br>
We need a mechanism to sent this interrupt from another task and receive a notification when the task interruption is complete.

## approach ##
Each node runs a task. <br>
Any node can try interrupting a task on another node. <br>
Interrupt is sent to all nodes via postgres notify. <br>
When this notification is received, the node running the task interupts it, and send a interrupt completion notification to all nodes. <br>
Any node competing with resources to run its own task can now do so. <br>
