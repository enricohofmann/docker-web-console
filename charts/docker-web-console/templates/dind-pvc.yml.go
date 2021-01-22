apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: {{ .Release.Name }}-dind
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 50Gi