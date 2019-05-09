package main

import (
	"github.com/msandim/web-hasher/hasher"
	"github.com/msandim/web-hasher/workerpool"
)

// Implementation of a workerpool job for the hashing case.
type hashJob struct {
	url    string
	hasher hasher.Hasher
	client hasher.HTTPClient
}

func (job hashJob) Process() workerpool.JobResult {
	hash, _ := job.hasher.Hash(job.url, job.client)
	return hashJobResult{job: job, hash: hash}
}

// Implementation of a workerpool job result for the hashing case.
type hashJobResult struct {
	job  hashJob
	hash string
}

func (result hashJobResult) GetJob() workerpool.Job {
	return result.job
}
