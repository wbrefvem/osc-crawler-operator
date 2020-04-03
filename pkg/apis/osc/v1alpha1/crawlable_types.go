package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CrawlableSpec defines the desired state of Crawlable
type CrawlableSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// URL is the thing to crawl
	URL string `json:"url"`
}

// CrawlableStatus defines the observed state of Crawlable
type CrawlableStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// CrawlJobs is a list of crawls for this Crawlable
	CrawlJobs []CrawlJob `json:"crawlJobs"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Crawlable is the Schema for the crawlables API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=crawlables,scope=Namespaced
type Crawlable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CrawlableSpec   `json:"spec,omitempty"`
	Status CrawlableStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CrawlableList contains a list of Crawlable
type CrawlableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Crawlable `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Crawlable{}, &CrawlableList{})
}
