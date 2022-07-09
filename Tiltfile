# For more on Extensions, see: https://docs.tilt.dev/extensions.html
# load('ext://restart_process', 'docker_build_with_restart')

redis = helm(
  'tilt/redis',
  name='redis-local',
  set=['usePassword=false','auth.enabled=false','cluster.enabled=false']
  )
k8s_yaml(redis)

etcd = helm(
  'tilt/etcd',
  name='etcd-local',
  set=['auth.rbac.enabled=false', 'cluster.enabled=false']
  )
k8s_yaml(etcd)

mysql = helm(
  'tilt/mysql',
  name='mysql-local',
  set=['image.tag=5.7', 'auth.rootPassword=zxcvbnm','readinessProbe.enabled=false']
  )
k8s_yaml(mysql)

k8s_resource('etcd-local', port_forwards=['2379:2379','2380:2380'])
k8s_resource('mysql-local', port_forwards=['3306:3306'])
k8s_resource('redis-local-master', port_forwards=['6379:6379'])