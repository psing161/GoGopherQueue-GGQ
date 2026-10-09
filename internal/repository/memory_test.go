package repository 
import ( 
	"testing" 

	"GoGopherQueue-GGQ/internal/job" 
	
)

func TestCreateAndGetJob(t * testing.T){
	repo := NewMemoryRepository()
	newJob := job.Job{
		ID: "Job-123",
		Type: "Test",
		Payload: "Test sample job",
		Status: "queued",
	}

	if err := repo.Create(newJob); err != nil{
		t.Fatalf("Create() returned an error %v", err)
	}

	got, err := repo.Get("Job-123")
	if err != nil{
		t.Fatalf("Get() returned an error %v", err)
	}

	if got.ID != newJob.ID{
		t.Errorf("got ID %q, want %q", got.ID, newJob.ID)
	}

	if got.Status != newJob.Status { t.Errorf("got status %q, want %q", got.Status, newJob.Status) }
}

func TestDeleteJob(t *testing.T) { 
	repo := NewMemoryRepository() 
	newJob := job.Job{ 
		ID: "job-123", 
		Type: "email", 
		Status: "queued", 
		} 
		
	if err := repo.Create(newJob); 
	err != nil { 
		t.Fatalf("Create() returned an error: %v", err) } 
	
	if err := repo.Delete("job-123"); 
	
	err != nil { 
		t.Fatalf("Delete() returned an error: %v", err) 
		} 
	_, err := repo.Get("job-123") 
	
	if err == nil { 
		t.Fatal("Get() after deletion should return an error") }
	
	} 
	
	
	func TestGetMissingJob(t *testing.T) { 
		repo := NewMemoryRepository() 
		_, err := repo.Get("missing-job") 
		if err == nil { 
			t.Fatal("Get() for a missing job should return an error") 
			} 
	}