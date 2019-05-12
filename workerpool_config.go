package main

import (
	"github.com/msandim/web-hasher/hasher"
	"github.com/msandim/web-hasher/workerpool"
)

// Implementation of a workerpool job for the hashing case.
type hashJob struct {
	url    string
	hasher hasher.Hasher
}

func (job hashJob) Process() workerpool.JobResult {
	hash, err := job.hasher.Hash(job.url)
	return hashJobResult{job: job, hash: hash, err: err}
}

// Implementation of a workerpool job result for the hashing case.
type hashJobResult struct {
	job  hashJob
	hash string
	err  error
}

func (result hashJobResult) GetJob() workerpool.Job {
	return result.job
}
