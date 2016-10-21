# traot

tool
requests
are 
off
topic

http://stackoverflow.com/questions/40160222/is-there-a-java-to-golang-converter-program

Run traot on a big java project like https://github.com/wildfly/wildfly and it will make a dir called `wildfly_go`
and break up all the endless com dot whatever dot whatever directories into a more flat structure like golang.
You can then cd into top level dirs and ls the files without having to cd in and out of a million dirs.
The goal is to help you read the java code in a more golang like way.

```
./traot "/Users/aa/dev/wildfly"

~/dev/wildfly_go $ ls -l
drwxr-xr-x     3 aa  staff     102 Oct 21 13:08 apache_directory
drwxr-xr-x    34 aa  staff    1156 Oct 21 13:08 hibernate_jpa
drwxr-xr-x  5837 aa  staff  198458 Oct 21 13:08 jboss_as
drwxr-xr-x     3 aa  staff     102 Oct 21 13:08 jboss_iiop
drwxr-xr-x     6 aa  staff     204 Oct 21 13:08 jboss_system
drwxr-xr-x     4 aa  staff     136 Oct 21 13:08 jipijapa_cache
drwxr-xr-x     8 aa  staff     272 Oct 21 13:08 jipijapa_eclipselink
drwxr-xr-x     5 aa  staff     170 Oct 21 13:08 jipijapa_event
drwxr-xr-x     7 aa  staff     238 Oct 21 13:08 jipijapa_management
drwxr-xr-x    12 aa  staff     408 Oct 21 13:08 jipijapa_plugin
drwxr-xr-x   503 aa  staff   17102 Oct 21 13:08 wildfly_clustering
drwxr-xr-x   715 aa  staff   24310 Oct 21 13:08 wildfly_extension
drwxr-xr-x   110 aa  staff    3740 Oct 21 13:08 wildfly_iiop
drwxr-xr-x     5 aa  staff     170 Oct 21 13:08 wildfly_jberet
drwxr-xr-x    25 aa  staff     850 Oct 21 13:08 wildfly_mod_cluster
drwxr-xr-x     8 aa  staff     272 Oct 21 13:08 wildfly_naming
drwxr-xr-x    62 aa  staff    2108 Oct 21 13:08 wildfly_test

~/dev/wildfly_go/wildfly_clustering $ ls -l ee_infinispan_*
-rwxr-xr-x  1 aa  staff  340 Oct 21 13:08 ee_infinispan_cacheentrymutator.go
-rwxr-xr-x  1 aa  staff  122 Oct 21 13:08 ee_infinispan_cacheentrymutatortestcase.go
-rwxr-xr-x  1 aa  staff   37 Oct 21 13:08 ee_infinispan_cacheproperties.go
-rwxr-xr-x  1 aa  staff   38 Oct 21 13:08 ee_infinispan_creator.go
-rwxr-xr-x  1 aa  staff   32 Oct 21 13:08 ee_infinispan_evictor.go
-rwxr-xr-x  1 aa  staff  371 Oct 21 13:08 ee_infinispan_infinispanbatch.go
-rwxr-xr-x  1 aa  staff  800 Oct 21 13:08 ee_infinispan_infinispanbatcher.go
-rwxr-xr-x  1 aa  staff  935 Oct 21 13:08 ee_infinispan_infinispanbatchertestcase.go
-rwxr-xr-x  1 aa  staff  489 Oct 21 13:08 ee_infinispan_infinispancacheproperties.go
-rwxr-xr-x  1 aa  staff  211 Oct 21 13:08 ee_infinispan_infinispancachepropertiestestcase.go
-rwxr-xr-x  1 aa  staff   35 Oct 21 13:08 ee_infinispan_locator.go
-rwxr-xr-x  1 aa  staff  207 Oct 21 13:08 ee_infinispan_mutablecacheentry.go
-rwxr-xr-x  1 aa  staff   54 Oct 21 13:08 ee_infinispan_mutator.go
-rwxr-xr-x  1 aa  staff   32 Oct 21 13:08 ee_infinispan_remover.go
-rwxr-xr-x  1 aa  staff  383 Oct 21 13:08 ee_infinispan_retryinginvoker.go
-rwxr-xr-x  1 aa  staff   52 Oct 21 13:08 ee_infinispan_transactionbatch.go

~/dev/wildfly_go/wildfly_clustering $ cat ee_infinispan_infinispanbatchertestcase.go
//public class InfinispanBatcherTestCase {
//private final TransactionManager tm = mock(TransactionManager.class);
//private final Batcher<TransactionBatch> batcher = new InfinispanBatcher(this.tm);
//public void destroy() {
//public void createExistingBatch() throws Exception {
//public void createBatchClose() throws Exception {
//public void createBatchDiscard() throws Exception {
//public void createNestedBatchClose() throws Exception {
//public void createNestedBatchDiscard() throws Exception {
//public void createOverlappingBatchClose() throws Exception {
//public void createOverlappingBatchDiscard() throws Exception {
//public void resumeNullBatch() throws Exception {
//public void resumeNonTxBatch() throws Exception {
//public void resumeBatch() throws Exception {
//public void resumeBatchExisting() throws Exception {
//public void suspendBatch() throws Exception {
//public void suspendNoBatch() throws Exception {
```
Future plans: write out more the java code methods and convert them to golang func's
