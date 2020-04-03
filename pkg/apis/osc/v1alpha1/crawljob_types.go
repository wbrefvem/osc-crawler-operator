package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CrawlJobSpec defines the desired state of CrawlJob
type CrawlJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// StartURLs tells the crawler where to start crawling
	StartURLs []string `json:"start-urls"`

	// AllowedDomains restricts the domains the crawler will crawl
	AllowedDomains []string `json:"allowedDomains"`
}

// CrawlJobStatus defines the observed state of CrawlJob
type CrawlJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Started is a timestamp that denotes when this crawl started
	TimeStarted string `json:"timeStarted"`

	// TimeOfLastStateChange is a timestamp updated every time State changes
	TimeOfLastStateChange string `json:"timeOfLastStateChange"`

	// State describes the current crawl (one of Pending, Succeeded, Failed)
	State string `json:"state"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CrawlJob is the Schema for the crawljobs API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=crawljobs,scope=Namespaced
type CrawlJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CrawlJobSpec   `json:"spec,omitempty"`
	Status CrawlJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CrawlJobList contains a list of CrawlJob
type CrawlJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CrawlJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CrawlJob{}, &CrawlJobList{})
}
